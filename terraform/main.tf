provider "digitalocean" {
    token = "${var.do_token}"
}

resource "digitalocean_droplet" "influxdb" {
    name = "influxdb"
    region = "tor1"
    size = "1gb"
    image = "coreos-stable"

    ssh_keys = [316099]

    connection {
        user = "core"
    }

    provisioner "file" {
        source = "files/influxdb/"
        destination = "/tmp/"
    }

    provisioner "remote-exec" {
        inline = [
            "sudo mkdir -p /var/opt/influxdb/",
            "sudo mkdir -p /etc/opt/influxdb/",
            "sudo mkdir -p /opt/influxdb/",

            "wget https://s3.amazonaws.com/influxdb/influxdb_0.9.4.2_x86_64.tar.gz",
            "tar xvfz influxdb_0.9.4.2_x86_64.tar.gz",
            "sudo mv influxdb_0.9.4.2_x86_64/opt/influxdb/versions/0.9.4.2/influxd /opt/influxdb/influxd",
            "sudo chmod +x /opt/influxdb/influxd",

            "sudo mv /tmp/influxdb.conf /etc/opt/influxdb/influxdb.conf",
            "sudo mv /tmp/influxdb.service /etc/systemd/system/influxdb.service",
            "sudo systemctl enable influxdb.service",
            "sudo systemctl start influxdb.service",
        ]
    }
}

resource "digitalocean_droplet" "scheduler" {
    name = "scheduler"
    region = "tor1"
    size = "1gb"
    image = "coreos-stable"

    ssh_keys = [316099]

    connection {
        user = "core"
    }

    provisioner "local-exec" {
        command = "GOOS=linux go build -o ./files/scheduler/schedulerd github.com/lgpeterson/loadtests/cmd/schedulerd"
    }

    provisioner "local-exec" {
        command = "GOOS=linux go build -o ./files/scheduler/executord github.com/lgpeterson/loadtests/executor/cmd/executord"
    }

    provisioner "file" {
        source = "files/scheduler/"
        destination = "/tmp/"
    }

    provisioner "remote-exec" {
        inline = [
            "sudo mkdir -p /opt/",
            "sudo mv /tmp/schedulerd /opt/schedulerd",
            "sudo mv /tmp/executord /opt/executord",
            "sudo chmod +x /opt/schedulerd",
            "sudo mkdir -p /etc/scheduler/",

            "sudo echo INFLUX_ADDR=${digitalocean_droplet.influxdb.ipv4_address}:${var.influx_port} >> /etc/scheduler/scheduler.env",
            "sudo echo INFLUX_DB_NAME=${var.influx_dbname} >> /etc/scheduler/scheduler.env",
            "sudo echo INFLUX_USERNAME=${var.influx_username} >> /etc/scheduler/scheduler.env",
            "sudo echo INFLUX_PASSWORD=${var.influx_password} >> /etc/scheduler/scheduler.env",
            "sudo echo DO_TOKEN=${var.do_token} >> /etc/scheduler/scheduler.env",

            "sudo mv /tmp/scheduler.env /etc/scheduler/scheduler.env",
            "sudo mv /tmp/scheduler.service /etc/systemd/system/scheduler.service",
            "sudo systemctl enable scheduler.service",
            "sudo systemctl start scheduler.service",
        ]
    }
}