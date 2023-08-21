packer {
  required_plugins {
    docker = {
      source  = "github.com/hashicorp/docker"
      version = "~> 1"
    }
    chef = {
      source  = "github.com/hashicorp/chef"
      version = "~> 1"
    }
  }
}

variable "source_image_tag" {
  type    = string
  default = ""
}

variable "generated_image_repository" {
  type    = string
  default = ""
}

variable "generated_image_tag" {
  type    = string
  default = ""
}

source "docker" "ubuntu_amd" {
  commit = true
  image  = "ubuntu:22.04"
  platform = "linux/amd64"
}

source "docker" "ubuntu_arm64" {
  commit = true
  image  = "ubuntu:22.04"
  platform = "linux/arm64/v8"
}

source "docker" "base_amd" {
  changes = ["ENV PATH /px/bin:/opt/google-cloud-sdk/bin:/opt/px_dev/bin:/opt/px_dev/tools/golang/bin:/opt/px_dev/tools/node/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin", "ENV GOPATH /px", "WORKDIR /px/src/px.dev/pixie", "ENTRYPOINT [\"\"]", "CMD [\"\"]"]
  commit  = true
  image   = "${var.generated_image_repository}/base_image/${var.source_image_tag}"
  platform = "linux/amd64"
}

source "docker" "base_arm64" {
  changes = ["ENV PATH /px/bin:/opt/google-cloud-sdk/bin:/opt/px_dev/bin:/opt/px_dev/tools/golang/bin:/opt/px_dev/tools/node/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin", "ENV GOPATH /px", "WORKDIR /px/src/px.dev/pixie", "ENTRYPOINT [\"\"]", "CMD [\"\"]"]
  commit  = true
  image   = "${var.generated_image_repository}/base_image/${var.source_image_tag}"
  platform = "linux/arm64/v8"
}

source "docker" "dev_amd" {
  changes = ["ENV PATH /px/bin:/opt/google-cloud-sdk/bin:/opt/px_dev/bin:/opt/px_dev/tools/golang/bin:/opt/px_dev/tools/node/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin", "ENV GOPATH /px", "WORKDIR /px/src/px.dev/pixie", "ENTRYPOINT [\"\"]", "CMD [\"\"]"]
  commit  = true
  image   = "${var.generated_image_repository}/dev_image/${var.source_image_tag}"
  platform = "linux/amd64"
}

source "docker" "dev_arm64" {
  changes = ["ENV PATH /px/bin:/opt/google-cloud-sdk/bin:/opt/px_dev/bin:/opt/px_dev/tools/golang/bin:/opt/px_dev/tools/node/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin", "ENV GOPATH /px", "WORKDIR /px/src/px.dev/pixie", "ENTRYPOINT [\"\"]", "CMD [\"\"]"]
  commit  = true
  image   = "${var.generated_image_repository}/dev_image/${var.source_image_tag}"
  platform = "linux/arm64/v8"
}

locals {
  run_list = {
    base_image = ["recipe[px_dev::linux]"]
    dev_image_with_extras = ["recipe[px_dev_extras]"]
    dev_image = ["role[px_base_dev]", "recipe[px_dev::docker_extras]", "recipe[px_dev::cleanup]"]
    linters_image = ["recipe[px_dev::setup]", "recipe[px_dev::linters]", "recipe[px_dev::docker_extras]", "recipe[px_dev::cleanup]"]
  }
}

build {
  name = "image"

  source "docker.ubuntu_amd" {
    name = "base_image_amd"
  }
  source "docker.ubuntu_arm64" {
    name = "base_image_arm64"
  }
  source "docker.base_amd" {
    name = "linters_image_amd"
  }
  source "docker.base_arm64" {
    name = "linters_image_arm64"
  }
  source "docker.base_amd" {
    name = "dev_image_amd"
  }
  source "docker.base_arm64" {
    name = "dev_image_arm64"
  }
  source "docker.dev_amd" {
    name = "dev_image_with_extras_amd"
  }
  source "docker.dev_arm64" {
    name = "dev_image_with_extras_arm64"
  }

  provisioner "shell" {
    inline = ["echo \"Building ${source.name}\""]
  }

  provisioner "shell" {
    inline = ["apt-get -y update; DEBIAN_FRONTEND=noninteractive apt-get install -y curl"]
  }

  provisioner "shell" {
    inline = ["DEBIAN_FRONTEND=noninteractive apt-get install -y git xz-utils"]
    only = ["source.docker.linter_image"]
  }

  provisioner "chef-solo" {
    cookbook_paths = ["cookbooks"]
    prevent_sudo   = true
    roles_path     = "roles"
    run_list       = local.run_list[trimsuffix(trimsuffix(source.name, "_amd"), "_arm64")]
    version        = "18.2.7"
  }

  provisioner "shell" {
    inline = ["rm -rf /opt/chef"]
  }

  post-processors {
    post-processor "docker-tag" {
      repository = "${var.generated_image_repository}/${source.name}"
      tags       = ["${var.generated_image_tag}"]
    }
    # post-processor "docker-push" {
    # }
  }
}
