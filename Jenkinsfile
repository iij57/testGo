pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        bat(script: 'cd src', returnStatus: true)
        bat 'cd main'
        bat 'go build .'
      }
    }

  }
}