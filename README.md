# kubernetes-basics


# Install docker 


# Install Kubernetes 
https://kubernetes.io/docs/setup/independent/install-kubeadm/


# Configure kubernetes 

kubeadm init --pod-network-cidr=192.168.0.0/16

kubectl apply -f https://docs.projectcalico.org/v3.3/getting-started/kubernetes/installation/hosted/rbac-kdd.yaml

kubectl apply -f https://docs.projectcalico.org/v3.3/getting-started/kubernetes/installation/hosted/kubernetes-datastore/calico-networking/1.7/calico.yaml



# check the pods

kubectl get pods --all-namespaces

# If you want to run everything(pods) in the same VM use the command
 
 kubectl taint nodes --all node-role.kubernetes.io/master-


# application deployments
clone the repo and go to "kubernetes-basics" directory and execute the files

kubectl create -f pvolume.yaml

kubectl create -f deployment.yaml

kubectl create -f autoscale.yaml
