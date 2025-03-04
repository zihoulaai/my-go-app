pipeline {
    agent any
    environment {
        DOCKER_IMAGE = "huzigege/my-go-app"
        DOCKER_CREDENTIALS = credentials('docker-hub-cred')
    }
    stages {
        stage('Checkout') {
            steps {
                git url: 'https://github.com/your-repo/my-go-app.git', branch: 'main'
            }
        }
        stage('Test') {
            steps {
                sh 'go test -v ./...'
            }
        }
        stage('Build Docker Image') {
            steps {
                script {
                    docker.build("${DOCKER_IMAGE}:${env.BUILD_NUMBER}")
                }
            }
        }
        stage('Push Image') {
            steps {
                script {
                    docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-cred') {
                        docker.image("${DOCKER_IMAGE}:${env.BUILD_NUMBER}").push()
                    }
                }
            }
        }
        stage('Deploy') {
            when {
                branch 'main'  # 仅 main 分支触发部署
            }
            steps {
                sh """
                    docker stop my-go-app || true
                    docker rm my-go-app || true
                    docker run -d --name my-go-app -p 8080:8080 ${DOCKER_IMAGE}:${env.BUILD_NUMBER}
                """
            }
        }
    }
    post {
        failure {
            emailext body: '构建失败，请检查日志: ${BUILD_URL}', subject: 'CI/CD Pipeline Failed'
        }
    }
}