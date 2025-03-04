pipeline {
    agent any
    
    environment {
        REPO_URL = 'https://github.com/zihoulaai/my-go-app.git'
        IMAGE_NAME = 'huzigege/my-go-app'
        IMAGE_TAG = 'latest'
        DOCKER_CREDENTIALS_ID = 'docker-hub-credentials'
    }
    
    stages {
        stage('Checkout') {
            steps {
                git branch: 'main', 
                    credentialsId: 'github-credentials', 
                    url: env.REPO_URL
            }
        }
        
        stage('Build Docker Image') {
            steps {
                script {
                    sh '/usr/local/bin/docker build -t $IMAGE_NAME:$IMAGE_TAG .'
                }
            }
        }
        
        stage('Login to Docker Hub') {
            steps {
                script {
                    withCredentials([string(credentialsId: 'docker-hub-credentials', variable: 'DOCKER_TOKEN')]) {
                        sh "echo $DOCKER_TOKEN | /usr/local/bin/docker login --username your-dockerhub-username --password-stdin"
                    }
                }
            }
        }
        
        stage('Push to Docker Hub') {
            steps {
                script {
                    sh '/usr/local/bin/docker push $IMAGE_NAME:$IMAGE_TAG'
                }
            }
        }
        
        stage('Cleanup') {
            steps {
                script {
                    sh '/usr/local/bin/docker rmi $IMAGE_NAME:$IMAGE_TAG'
                }
            }
        }
    }
}

// pipeline {
//     agent any

//     environment {
//         GOROOT='/usr/local/opt/go/libexec'
//         PATH="${GOROOT}/bin:${env.PATH}"
//         IMAGE_NAME="my-go-app"
//         IMAGE_TAG="latest"
//         REGISTRY="docker.io/huzigege"
//     }

//     options {
//         timeout(time: 30, unit: 'MINUTES') // 限制 CI/CD 超时时间
//         disableConcurrentBuilds() // 禁止并发执行
//     }

//     stages {
//         stage('Checkout Code') {
//             steps {
//                 git branch: 'main', 
//                     credentialsId: 'github-credentials', 
//                     url: 'https://github.com/zihoulaai/my-go-app.git'
//             }
//         }

//         stage('Cache Dependencies') {
//             steps {
//                 sh 'go mod download'
//             }
//         }

//         // stage('Build') {
//         //     steps {
//         //         sh 'go build -o app'
//         //     }
//         // }

//         stage('Test & Static Analysis') {
//             parallel {
//                 stage('Unit Tests') {
//                     steps {
//                         sh 'go test ./... -v'
//                     }
//                 }
//                 // stage('Static Code Analysis') {
//                 //     steps {
//                 //         sh 'golangci-lint run || true' // 允许 lint 失败但不中断 CI
//                 //     }
//                 // }
//             }
//         }

//         stage('Build Docker Image') {
//             steps {
//                 sh """
//                 /usr/local/bin/docker build -t ${REGISTRY}/${IMAGE_NAME}:${IMAGE_TAG} .
//                 """
//             }
//         }

//         stage('Push Docker Image') {
//             steps {
//                 withDockerRegistry([credentialsId: 'docker-hub-credentials', url: '']) {
//                     sh "/usr/local/bin/docker push ${REGISTRY}/${IMAGE_NAME}:${IMAGE_TAG}"
//                 }
//             }
//         }

//         // stage('Deploy to Kubernetes') {
//         //     steps {
//         //         withKubeConfig([credentialsId: 'k8s-credentials']) {
//         //             sh """
//         //             kubectl set image deployment/${IMAGE_NAME} ${IMAGE_NAME}=${REGISTRY}/${IMAGE_NAME}:${IMAGE_TAG}
//         //             """
//         //         }
//         //     }
//         // }
//     }

//     post {
//         success {
//             echo "CI/CD Pipeline executed successfully!"
//         }
//         failure {
//             echo "Pipeline failed!"
//         }
//     }
// }