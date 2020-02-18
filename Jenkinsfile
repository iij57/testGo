pipeline {
  agent {
    node {
      label 'go_build'
    }

  }
  stages {
    stage('Build') {
      steps {
        bat 'cd src'
        bat 'cd main'
        bat 'go build .'
      }
    }

  }
}