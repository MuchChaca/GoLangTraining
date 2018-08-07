# Trump face finder

## Create the pipeline inputs
1. ``training`` - which includes images of faces that we use to “teach” facebox
2. ``unidentified`` - which includes images with faces we want to detect and identify
3. ``labels`` - which includes label images that we will overlay on the unidentified images to indicate identified faces

**Example using ``pachctl``:**
```bash
# Start docker
➔ systemctl start docker
# Start minikube
➔ minikube start 
# Configure pods
➔ pachctl deploy local
# Set up port forwarding so commands you send can reach Pachyderm within the VM. We background this process since port forwarding blocks.
➔ pachctl port-forward &

➔ pachctl create-repo training
➔ pachctl create-repo unidentified
➔ pachctl create-repo labels
➔ pachctl list-repo
NAME                CREATED             SIZE                
labels              3 seconds ago       0 B                 
unidentified        11 seconds ago      0 B                 
training            17 seconds ago      0 B
➔ cd data/train/faces1/
➔ ls
trump1.jpg  trump2.jpg  trump3.jpg  trump4.jpg  trump5.jpg
➔ pachctl put-file training master -c -r -f . # TODO: detail that
➔ pachctl list-repo
NAME                CREATED             SIZE                
training            5 minutes ago       486.2 KiB           
labels              5 minutes ago       0 B                 
unidentified        5 minutes ago       0 B                 
➔ pachctl list-file training master
NAME                TYPE                SIZE                
trump1.jpg          file                78.98 KiB           
trump2.jpg          file                334.5 KiB           
trump3.jpg          file                11.63 KiB           
trump4.jpg          file                27.45 KiB           
trump5.jpg          file                33.6 KiB
➔ cd ../../labels/
➔ ls
clinton.jpg  trump.jpg
➔ pachctl put-file labels master -c -r -f .
➔ cd ../unidentified/
➔ ls
image1.jpg  image2.jpg
➔ pachctl put-file unidentified master -c -r -f .
➔ pachctl list-repo
NAME                CREATED             SIZE                
unidentified        7 minutes ago       540.4 KiB           
labels              7 minutes ago       15.44 KiB           
training            7 minutes ago       486.2 KiB
```


## Create the pipelines

### Pipeline: ``train.json``
``train.json``
```json
{
  "pipeline": {
    "name": "model"
  },
  "transform": {

    // we use the facebox docker image (from machinebox)
    "image": "dwhitena/fbtrain",

    // what command to use in the docker container
    "cmd": [ "/bin/bash" ],

    // data to input to the pipeline, then persist the state of the model
    "stdin": [
	"./facebox &>/dev/null </dev/zero &",
	"sleep 5",
	"/gotrain -inDir=/pfs/training -outDir=/pfs/out"
    ],

    "env": {
	"MB_KEY": "NWFiYTJlMWNkNWJlY2U5NzVlMmIxN2FlY2IzMjIyYjg.4YKkENp7h71YtdsimmD0lzrhoWRtmR19aXbJi3Y2sJ9YiqBvjLLk4psPxt2Ct8Wl3AjsPel0TX9Txcw6yc24jw"
    }
  },
  "parallelism_spec": {
    "constant": "1"
  },
  "input": {
    "atom": {
      "repo": "training",
      "glob": "/"
    }
  }
}
```

```bash
create-MB-pipeline.sh  identify.json  tag.json  train.json
➔ ./create-MB-pipeline.sh train.json
➔ pachctl list-pipeline
NAME                INPUT               OUTPUT              STATE               
model               training            model/master        running    
➔ pachctl list-job
ID                                   OUTPUT COMMIT STARTED       DURATION RESTART PROGRESS STATE            
3425a7a0-543e-4e2a-a244-a3982c527248 model/-       9 seconds ago -        1       0 / 1    running
➔ pachctl list-job
ID                                   OUTPUT COMMIT                          STARTED       DURATION  RESTART PROGRESS STATE            
3425a7a0-543e-4e2a-a244-a3982c527248 model/1b9c158e33394056a18041a4a86cb54a 5 minutes ago 5 minutes 1       1 / 1    success
➔ pachctl list-repo
NAME                CREATED             SIZE                
model               5 minutes ago       4.118 KiB           
unidentified        18 minutes ago      540.4 KiB           
labels              18 minutes ago      15.44 KiB           
training            19 minutes ago      486.2 KiB           
➔ pachctl list-file model master
NAME                TYPE                SIZE                
state.facebox       file                4.118 KiB
```

## Use the trained facebox to identify faces
We make a new pipeline based on a ``identify.json`` to identify faces of the ``unidentified`` pipeline. This will take the persisted state of the model.

> It will also execute cURL commands to interact with facebox, and it will output indications of identified faces to JSON files, one per ``unidentified`` image.

```bash
➔ ./create-MB-pipeline.sh identify.json
➔ pachctl list-job
ID                                   OUTPUT COMMIT                          STARTED       DURATION  RESTART PROGRESS STATE            
281d4393-05c8-44bf-b5de-231cea0fc022 identify/-                             6 seconds ago -         0       0 / 2    running
3425a7a0-543e-4e2a-a244-a3982c527248 model/1b9c158e33394056a18041a4a86cb54a 8 minutes ago 5 minutes 1       1 / 1    success
➔ pachctl list-job
ID                                   OUTPUT COMMIT                             STARTED            DURATION   RESTART PROGRESS STATE            
281d4393-05c8-44bf-b5de-231cea0fc022 identify/287fc78a4cdf42d89142d46fb5f689d9 About a minute ago 53 seconds 0       2 / 2    success
3425a7a0-543e-4e2a-a244-a3982c527248 model/1b9c158e33394056a18041a4a86cb54a    9 minutes ago      5 minutes  1       1 / 1    success
➔ pachctl list-repo
NAME                CREATED              SIZE                
identify            About a minute ago   1.932 KiB           
model               10 minutes ago       4.118 KiB           
unidentified        23 minutes ago       540.4 KiB           
labels              23 minutes ago       15.44 KiB           
training            24 minutes ago       486.2 KiB           
➔ pachctl list-file identify master
NAME                TYPE                SIZE                
image1.json         file                1.593 KiB           
image2.json         file                347 B
```

## Tag identified faces

```bash
➔ pachctl create-pipeline -f tag.json
➔ pachctl list-job
ID                                   OUTPUT COMMIT                             STARTED        DURATION   RESTART PROGRESS STATE            
cd284a28-6c97-4236-9f6d-717346c60f24 tag/-                                     2 seconds ago  -          0       0 / 2    running
281d4393-05c8-44bf-b5de-231cea0fc022 identify/287fc78a4cdf42d89142d46fb5f689d9 5 minutes ago  53 seconds 0       2 / 2    success
3425a7a0-543e-4e2a-a244-a3982c527248 model/1b9c158e33394056a18041a4a86cb54a    13 minutes ago 5 minutes  1       1 / 1    success
➔ pachctl list-job
ID                                   OUTPUT COMMIT                             STARTED        DURATION   RESTART PROGRESS STATE            
cd284a28-6c97-4236-9f6d-717346c60f24 tag/ae747e8032704b6cae6ae7bba064c3c3      25 seconds ago 11 seconds 0       2 / 2    success
281d4393-05c8-44bf-b5de-231cea0fc022 identify/287fc78a4cdf42d89142d46fb5f689d9 5 minutes ago  53 seconds 0       2 / 2    success
3425a7a0-543e-4e2a-a244-a3982c527248 model/1b9c158e33394056a18041a4a86cb54a    14 minutes ago 5 minutes  1       1 / 1    success
➔ pachctl list-repo
NAME                CREATED             SIZE                
tag                 30 seconds ago      591.3 KiB           
identify            5 minutes ago       1.932 KiB           
model               14 minutes ago      4.118 KiB           
unidentified        27 minutes ago      540.4 KiB           
labels              27 minutes ago      15.44 KiB           
training            27 minutes ago      486.2 KiB           
➔ pachctl list-file tag master
NAME                TYPE                SIZE                
tagged_image1.jpg   file                557 KiB             
tagged_image2.jpg   file                34.35 KiB  
```