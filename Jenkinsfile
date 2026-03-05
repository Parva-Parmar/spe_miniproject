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
                    // Checkout the code from the GitHub repository [cite: 170]
                    git branch: 'main', url: "${GITHUB_REPO_URL}"
                }
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    // Build Docker image [cite: 180]
                    // We tag it directly with your username to make pushing easier
                    docker.build("${DOCKERHUB_USERNAME}/${DOCKER_IMAGE_NAME}:latest", '.')
                }
            }
        }

        stage('Push Docker Images') {
            steps {
                script{
                    // Push to Docker Hub using credentials we will configure in Jenkins later [cite: 198]
                    docker.withRegistry('', 'DockerHubCred') {
                        sh "docker push ${DOCKERHUB_USERNAME}/${DOCKER_IMAGE_NAME}:latest"
                    }
                }
            }
        }

        stage('Run Ansible Playbook') {
            steps {
                script {
                    // Executes the Ansible playbook using our inventory file [cite: 210]
                    ansiblePlaybook(
                        playbook: 'deploy.yml',
                        inventory: 'inventory'
                    )
                }
            }
        }
    }
}