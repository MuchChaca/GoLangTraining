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
➔ pachctl put-file training master -c -r -f .
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

