
tag = true

pipeline "build" {
  step {
    type = "docker"
    image = "rls"
    context = "."
  }

  step {
    type = "docker-shell"
    image = "rls"
    volume "code" {
      source = "."
      dest = "/go/src/github.com/altipla-consulting/rls"
    }
    volume "binary" {
      source = "./tmp"
      dest = "/opt/build"
    }
  }
}

pipeline "deploy" {
  step {
    type = "storage"
    bucket = "altipla-artifacts"
    path = "rls"
    public = true
    cache = false
    files = [
      "tmp/rls.${TAG}",
    ]
  }

  step {
    type = "git-tag"
    name = "${TAG}"
  }
}
