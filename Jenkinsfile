pipeline {
    agent any

    environment {
        DOCKERHUB_USERNAME = 'parva04'
        DOCKER_IMAGE_NAME = 'calculator'
        GITHUB_REPO_URL = 'https://github.com/Parva-Parmar/spe_miniproject.git'
        NOTIFY_EMAIL = 'parvaparmar41@gmail.com'
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
            // The post block is now correctly INSIDE the 'Test Code' stage
            post {
                success {
                    mail to: "${NOTIFY_EMAIL}",
                         subject: "Tests Passed: ${currentBuild.fullDisplayName}",
                         body: "Good news!\n\nThe Go unit tests for the Scientific Calculator passed successfully. The pipeline is now proceeding to build the Docker image.\n\nView logs: ${env.BUILD_URL}"
                }
                failure {
                    mail to: "${NOTIFY_EMAIL}",
                         subject: "Tests FAILED: ${currentBuild.fullDisplayName}",
                         body: "Attention:\n\nThe Go unit tests failed. The pipeline has been halted to prevent a broken image from being built.\n\nPlease check the logs to fix the code: ${env.BUILD_URL}"
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
                    sh "ansible all -m ping -i inventory.ini"
                }
            }
        }

        stage('Run Ansible Playbook') {
            steps {
                script {
                    ansiblePlaybook(
                        playbook: 'deploy.yml',
                        inventory: 'inventory.ini'
                    )
                }
            }
        }
    }


    post {
        success {
            mail to: "${NOTIFY_EMAIL}",
                 subject: "Deployment SUCCESS: ${currentBuild.fullDisplayName}",
                 body: "The Scientific Calculator pipeline has successfully completed all stages. The application is now deployed and running via Ansible.\n\nView pipeline details: ${env.BUILD_URL}"
        }
        failure {
            mail to: "${NOTIFY_EMAIL}",
                 subject: "Pipeline FAILED: ${currentBuild.fullDisplayName}",
                 body: "The pipeline encountered an error after the testing phase.\n\nPlease review the logs here: ${env.BUILD_URL}"
        }
    }
}