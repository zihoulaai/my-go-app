pipeline {
    agent any

    environment {
        GOROOT='/usr/local/opt/go/libexec'
        PATH="${GOROOT}/bin:${env.PATH}"
        IMAGE_NAME="my-go-app"
        IMAGE_TAG="latest"
        REGISTRY="docker.io/huzigege"
        DOCKER_HOST=unix:///Users/$(wuliuqi)/.docker/run/docker.sock
    }

    options {
        timeout(time: 30, unit: 'MINUTES') // 限制 CI/CD 超时时间
        disableConcurrentBuilds() // 禁止并发执行
    }

    stages {
        stage('Checkout Code') {
            steps {
                git branch: 'main', 
                    credentialsId: 'github-credentials', 
                    url: 'https://github.com/zihoulaai/my-go-app.git'
            }
        }

        stage('Cache Dependencies') {
            steps {
                sh 'go mod download'
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o app'
            }
        }

        stage('Test & Static Analysis') {
            parallel {
                stage('Unit Tests') {
                    steps {
                        sh 'go test ./... -v'
                    }
                }
                // stage('Static Code Analysis') {
                //     steps {
                //         sh 'golangci-lint run || true' // 允许 lint 失败但不中断 CI
                //     }
                // }
            }
        }

        stage('Build Docker Image') {
            steps {
                sh """
                docker build -t ${REGISTRY}/${IMAGE_NAME}:${IMAGE_TAG} .
                """
            }
        }

        stage('Push Docker Image') {
            steps {
                withDockerRegistry([credentialsId: 'docker-hub-credentials', url: '']) {
                    sh "docker push ${REGISTRY}/${IMAGE_NAME}:${IMAGE_TAG}"
                }
            }
        }

        // stage('Deploy to Kubernetes') {
        //     steps {
        //         withKubeConfig([credentialsId: 'k8s-credentials']) {
        //             sh """
        //             kubectl set image deployment/${IMAGE_NAME} ${IMAGE_NAME}=${REGISTRY}/${IMAGE_NAME}:${IMAGE_TAG}
        //             """
        //         }
        //     }
        // }
    }

    post {
        success {
            echo "CI/CD Pipeline executed successfully!"
        }
        failure {
            echo "Pipeline failed!"
        }
    }
}