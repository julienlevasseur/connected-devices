job "plugs" {
    datacenters = ["dc1"]

    type = "batch"

    periodic {
        cron = "0 */1 * * * * *"
        time_zone = "America/Toronto"
    }

    parameterized {
		payload       = "forbidden"
		meta_required = ["DURATION"]
		meta_required = ["ID"]
	}

    task "plug" {
        //driver = "exec"
        driver = "raw_exec"

        config {
			command = "/bin/bash"
    		args    = [
                "-c",
                "/Users/julien.levasseur/go/src/github.com/julienlevasseur/connected-devices/plugs/plugs",
			]
		}

        env {
            CD_PLUS_ID = "${NOMAD_META_ID}"
            CD_PLUGS_DURATION = "${NOMAD_META_DURATION}"
        }

        //artifact {
        //    mode = "file"
        //    source = "https://jlevasseur-artifact-storage.s3.amazonaws.com/julienlevasseur/plugs"
      	//}
    }
}