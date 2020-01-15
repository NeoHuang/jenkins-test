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
  triggers {
    GenericTrigger(
      token: 'abc123',
      causeString: 'Triggered on PR change and reviews'
      )
  }
  environment {
    GOCACHE = '/tmp'
  }
}
