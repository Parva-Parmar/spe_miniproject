pipeline {
    agent any

    environment {
        DOCKERHUB_USERNAME = 'parva04'
        DOCKER_IMAGE_NAME = 'calculator'
        GITHUB_REPO_URL = 'https://github.com/Parva-Parmar/spe_miniproject.git'
    }

    stages {
        stage('Checkout') {
            steps {
                script {
                    git branch: 'main', url: "${GITHUB_REPO_URL}"
                }
            }
        }

        stage('Test Code') {
            steps {
                script {
                    echo "Running unit tests..."
                    sh "docker run --rm -v \$(pwd):/app -w /app golang:1.22-alpine go test -v ./..."
                }
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    docker.build("${DOCKERHUB_USERNAME}/${DOCKER_IMAGE_NAME}:latest", '.')
                }
            }
        }

        stage('Push Docker Images') {
            steps {
                script{
                    docker.withRegistry('', 'DockerHubCred') {
                        sh "docker push ${DOCKERHUB_USERNAME}/${DOCKER_IMAGE_NAME}:latest"
                    }
                }
            }
        }
        stage('Ping Server') {
            steps {
                script {
                    echo "Checking if target deployment server is reachable..."
                    sh "ansible all -m ping -i inventory"
                }
            }
        }
        stage('Run Ansible Playbook') {
            steps {
                script {
                    ansiblePlaybook(
                        playbook: 'deploy.yml',
                        inventory: 'inventory'
                    )
                }
            }
        }
    }
}