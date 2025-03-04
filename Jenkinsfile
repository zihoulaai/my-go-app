pipeline {
    agent {
        docker {
            image 'golang:1.23'
        }
    }

    environment {
        GO111MODULE = 'on'  // 启用 Go Modules
        GOPROXY = 'https://proxy.golang.org,direct' // 代理加速下载
    }

    stages {
        stage('Checkout Code') {
            steps {
                git branch: 'main', 
                    credentialsId: 'github-credentials', 
                    url: 'https://github.com/zihoulaai/my-go-app.git'
            }
        }

        stage('Setup Go') {
            steps {
                script {
                    def goVersion = "1.20"  // 选择 Golang 版本
                    sh "wget https://go.dev/dl/go${goVersion}.linux-amd64.tar.gz -O go.tar.gz"
                    sh "sudo tar -C /usr/local -xzf go.tar.gz"
                    sh "export PATH=\$PATH:/usr/local/go/bin"
                    sh "go version"
                }
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

    post {
        success {
            echo 'Build completed successfully!'
        }
        failure {
            echo 'Build failed!'
        }
    }
}
