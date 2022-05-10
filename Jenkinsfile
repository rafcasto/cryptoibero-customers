pipeline {
    agent none
     environment {
        registry = "rafcasto/cryptoibero-customers"
        registryCredential = 'rafcasto-dockerhub-crls'
        dockerImage = ''
       
    }
    stages { 
        stage('docker-package'){
            agent {
                docker { 
                    image 'rafcasto/nodejs-build'
                    args '-v /var/run/docker.sock:/var/run/docker.sock:rw -v /usr/bin/docker:/usr/bin/docker:rw -u root:root'    
                }
            }
            steps {
                script {
                    dockerImage = docker.build registry
                    docker.withRegistry( '', registryCredential ) {
                        dockerImage.push()
                    }
                }
            }
        }
        stage('kub8-deployment'){
            agent {
                docker { 
                    image 'rafcasto/nodejs-build'
                    args '-v /usr/local/bin/kubectl:/usr/local/bin/kubectl  -v /root/.kube/config:/home/node/.kube/config -u root:root'    
                }
            }
            steps{
                 catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE'){
                   sh 'kubectl delete svc customers-svc -n cryptoibero --kubeconfig=/home/node/.kube/config'
                   sh 'kubectl delete -n cryptoibero  deployment customers-dep --kubeconfig=/home/node/.kube/config'
                 }
                sh 'kubectl apply -f customers-deployment.yaml -n cryptoibero --kubeconfig=/home/node/.kube/config'
                sh 'kubectl apply -f customers-service.yaml -n cryptoibero --kubeconfig=/home/node/.kube/config'
            }
        }
    }
   
}