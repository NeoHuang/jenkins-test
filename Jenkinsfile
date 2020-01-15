pipeline {
  agent {
    docker {
      image 'golang:1.12-stretch'
    }

  }
  stages {
    stage('build') {
      steps {
        sh 'go build'
      }
    }

    stage('test') {
      steps {
        sh 'go test'
      }
    }

  }
  environment {
    GOCACHE = '/tmp'
  }
}