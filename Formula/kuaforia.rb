class Kuaforia < Formula
  desc "CLI for Kuaforia case management"
  homepage "https://kuaforia.com"
  url "https://github.com/kuaforia/cli/releases/download/v1.0.0/kuaforia-darwin-amd64.tar.gz"
  sha256 "0000000000000000000000000000000000000000000000000000000000000000"
  version "1.0.0"

  def install
    bin.install "kuaforia"
  end
end
