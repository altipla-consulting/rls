
tag = true

pipeline "build" {
  step {
    type = "golang"
    project = "github.com/altipla-consulting/rls"
    package = "cmd/rls"
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
