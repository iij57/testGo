pipeline {
  agent {
    docker {
      image 'golang'
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