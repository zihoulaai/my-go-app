pipeline {
    agent any

    environment {
        REPO_URL = 'https://github.com/zihoulaai/my-go-app.git'
        IMAGE_NAME = 'huzigege/my-go-app'
        IMAGE_TAG = 'latest'
        DOCKER_CREDENTIALS_ID = 'docker-hub-credentials'  // 这里使用 Token
        DOCKER_USER = 'huzigege'  // 你的 Docker Hub 用户名
    }

    stages {
        stage('Checkout') {
            steps {
                checkout([
                    $class: 'GitSCM',
                    branches: [[name: '*/main']],
                    doGenerateSubmoduleConfigurations: false,
                    extensions: [[$class: 'WipeWorkspace']],  
                    submoduleCfg: [],
                    userRemoteConfigs: [[
                        url: env.REPO_URL,
                        credentialsId: 'github-credentials'
                    ]]
                ])
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    try {
                        sh '/usr/local/bin/docker build -t $IMAGE_NAME:$IMAGE_TAG .'
                    } catch (Exception e) {
                        error "Docker Build Failed: ${e.message}"
                    }
                }
            }
        }

        stage('Push to Docker Hub') {
            steps {
                withCredentials([string(credentialsId: env.DOCKER_CREDENTIALS_ID, variable: 'DOCKER_TOKEN')]) {
                    sh '''
                        echo "$DOCKER_TOKEN" | docker login -u "$DOCKER_USER" --password-stdin
                        docker push $IMAGE_NAME:$IMAGE_TAG
                        docker logout
                    '''
                }
            }
        }

        stage('Cleanup') {
            steps {
                script {
                    sh 'docker rmi $IMAGE_NAME:$IMAGE_TAG'
                }
            }
        }
    }

    post {
        always {
            sh 'docker logout'
            cleanWs()
        }
        failure {
            echo 'Build failed, please check logs!'
        }
    }
}