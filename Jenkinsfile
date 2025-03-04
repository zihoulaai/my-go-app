pipeline {
    agent {
        docker {
            image 'golang:1.20'
        }
    }
    environment {
        PATH = "/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/sbin"
    }
    stages {
        stage('Checkout Code') {
            steps {
                git branch: 'main', 
                    credentialsId: 'github-credentials', 
                    url: 'https://github.com/zihoulaai/my-go-app.git'
            }
        }

        stage('Install Dependencies') {
            steps {
                sh 'go mod tidy'
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o app'
            }
        }

        stage('Test') {
            steps {
                sh 'go test ./... -v'
            }
        }

        stage('Publish Artifact') {
            steps {
                archiveArtifacts artifacts: 'app', fingerprint: true
            }
        }
    }
}
