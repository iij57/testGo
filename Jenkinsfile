pipeline {
  agent {
    docker {
      image 'golang:windowsservercore-1809'
    }

  }
  stages {
    stage('Build') {
      steps {
        sh 'go get -u github.com/julienschmidt/httprouter'
        sh 'go build .'
      }
    }

  }
}