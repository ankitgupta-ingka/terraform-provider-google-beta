package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccComputeRegionUrlMap_update_path_matcher(t *testing.T) {
	t.Parallel()

	randomSuffix := acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeUrlMapDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionUrlMap_basic1(randomSuffix),
			},
			{
				ResourceName:      "google_compute_region_url_map.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionUrlMap_basic2(randomSuffix),
			},
			{
				ResourceName:      "google_compute_region_url_map.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRegionUrlMap_advanced(t *testing.T) {
	t.Parallel()

	randomSuffix := acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeUrlMapDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionUrlMap_advanced1(randomSuffix),
			},
			{
				ResourceName:      "google_compute_region_url_map.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionUrlMap_advanced2(randomSuffix),
			},
			{
				ResourceName:      "google_compute_region_url_map.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRegionUrlMap_noPathRulesWithUpdate(t *testing.T) {
	t.Parallel()

	randomSuffix := acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeUrlMapDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionUrlMap_noPathRules(randomSuffix),
			},
			{
				ResourceName:      "google_compute_region_url_map.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionUrlMap_basic1(randomSuffix),
			},
			{
				ResourceName:      "google_compute_region_url_map.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRegionUrlMap_ilbUpdate(t *testing.T) {
	t.Parallel()

	randomSuffix := acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeUrlMapDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionUrlMap_ilb(randomSuffix),
			},
			{
				ResourceName:      "google_compute_region_url_map.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionUrlMap_ilbUpdate(randomSuffix),
			},
			{
				ResourceName:      "google_compute_region_url_map.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeRegionUrlMap_basic1(randomSuffix string) string {
	return fmt.Sprintf(`
resource "google_compute_region_backend_service" "foobar" {
	region        = "us-central1"
	name          = "regionurlmap-test-%s"
	protocol      = "HTTP"
	health_checks = ["${google_compute_region_health_check.zero.self_link}"]
}

resource "google_compute_region_health_check" "zero" {
	region = "us-central1"
	name   = "regionurlmap-test-%s"
	http_health_check {
	}
}

resource "google_compute_region_url_map" "foobar" {
	region          = "us-central1"
	name            = "regionurlmap-test-%s"
	default_service = "${google_compute_region_backend_service.foobar.self_link}"

	host_rule {
		hosts        = ["mysite.com", "myothersite.com"]
		path_matcher = "boop"
	}

	path_matcher {
		default_service = "${google_compute_region_backend_service.foobar.self_link}"
		name            = "boop"

		path_rule {
			paths   = ["/*"]
			service = "${google_compute_region_backend_service.foobar.self_link}"
		}
	}

	test {
		host    = "mysite.com"
		path    = "/*"
		service = "${google_compute_region_backend_service.foobar.self_link}"
	}
}
`, randomSuffix, randomSuffix, randomSuffix)
}

func testAccComputeRegionUrlMap_basic2(randomSuffix string) string {
	return fmt.Sprintf(`
resource "google_compute_region_backend_service" "foobar" {
	region        = "us-central1"
	name          = "regionurlmap-test-%s"
	protocol      = "HTTP"
	health_checks = ["${google_compute_region_health_check.zero.self_link}"]
}

resource "google_compute_region_health_check" "zero" {
	region = "us-central1"
	name   = "regionurlmap-test-%s"
	http_health_check {
	}
}

resource "google_compute_region_url_map" "foobar" {
	region          = "us-central1"
	name            = "regionurlmap-test-%s"
	default_service = "${google_compute_region_backend_service.foobar.self_link}"

	host_rule {
		hosts        = ["mysite.com", "myothersite.com"]
		path_matcher = "blip"
	}

	path_matcher {
		default_service = "${google_compute_region_backend_service.foobar.self_link}"
		name            = "blip"

		path_rule {
			paths   = ["/*", "/home"]
			service = "${google_compute_region_backend_service.foobar.self_link}"
		}
	}

	test {
		host    = "mysite.com"
		path    = "/test"
		service = "${google_compute_region_backend_service.foobar.self_link}"
	}
}
`, randomSuffix, randomSuffix, randomSuffix)
}

func testAccComputeRegionUrlMap_advanced1(randomSuffix string) string {
	return fmt.Sprintf(`
resource "google_compute_region_backend_service" "foobar" {
	region        = "us-central1"
	name          = "regionurlmap-test-%s"
	protocol      = "HTTP"
	health_checks = ["${google_compute_region_health_check.zero.self_link}"]
}

resource "google_compute_region_health_check" "zero" {
	region = "us-central1"
	name   = "regionurlmap-test-%s"
	http_health_check {
	}
}

resource "google_compute_region_url_map" "foobar" {
	region          = "us-central1"
	name            = "regionurlmap-test-%s"
	default_service = "${google_compute_region_backend_service.foobar.self_link}"

	host_rule {
		hosts        = ["mysite.com", "myothersite.com"]
		path_matcher = "blop"
	}

	host_rule {
		hosts        = ["myfavoritesite.com"]
		path_matcher = "blip"
	}

	path_matcher {
		default_service = "${google_compute_region_backend_service.foobar.self_link}"
		name            = "blop"

		path_rule {
			paths   = ["/*", "/home"]
			service = "${google_compute_region_backend_service.foobar.self_link}"
		}
	}

	path_matcher {
		default_service = "${google_compute_region_backend_service.foobar.self_link}"
		name            = "blip"

		path_rule {
			paths   = ["/*", "/home"]
			service = "${google_compute_region_backend_service.foobar.self_link}"
		}
	}
}
`, randomSuffix, randomSuffix, randomSuffix)
}

func testAccComputeRegionUrlMap_advanced2(randomSuffix string) string {
	return fmt.Sprintf(`
resource "google_compute_region_backend_service" "foobar" {
	region        = "us-central1"
	name          = "regionurlmap-test-%s"
	protocol      = "HTTP"
	health_checks = ["${google_compute_region_health_check.zero.self_link}"]
}

resource "google_compute_region_health_check" "zero" {
	region = "us-central1"
	name = "regionurlmap-test-%s"
	http_health_check {
	}
}

resource "google_compute_region_url_map" "foobar" {
	region          = "us-central1"
	name            = "regionurlmap-test-%s"
	default_service = "${google_compute_region_backend_service.foobar.self_link}"

	host_rule {
		hosts        = ["mysite.com", "myothersite.com"]
		path_matcher = "blep"
	}

	host_rule {
		hosts        = ["myfavoritesite.com"]
		path_matcher = "blip"
	}

	host_rule {
		hosts        = ["myleastfavoritesite.com"]
		path_matcher = "blub"
	}

	path_matcher {
		default_service = "${google_compute_region_backend_service.foobar.self_link}"
		name            = "blep"

		path_rule {
			paths   = ["/home"]
			service = "${google_compute_region_backend_service.foobar.self_link}"
		}

		path_rule {
			paths   = ["/login"]
			service = "${google_compute_region_backend_service.foobar.self_link}"
		}
	}

	path_matcher {
		default_service = "${google_compute_region_backend_service.foobar.self_link}"
		name            = "blub"

		path_rule {
			paths   = ["/*", "/blub"]
			service = "${google_compute_region_backend_service.foobar.self_link}"
		}
	}

	path_matcher {
		default_service = "${google_compute_region_backend_service.foobar.self_link}"
		name            = "blip"

		path_rule {
			paths   = ["/*", "/home"]
			service = "${google_compute_region_backend_service.foobar.self_link}"
		}
	}
}
`, randomSuffix, randomSuffix, randomSuffix)
}

func testAccComputeRegionUrlMap_noPathRules(randomSuffix string) string {
	return fmt.Sprintf(`
resource "google_compute_region_backend_service" "foobar" {
	region        = "us-central1"
	name          = "regionurlmap-test-%s"
	protocol      = "HTTP"
	health_checks = ["${google_compute_region_health_check.zero.self_link}"]
}

resource "google_compute_region_health_check" "zero" {
	region = "us-central1"
	name   = "regionurlmap-test-%s"
	http_health_check {
	}
}

resource "google_compute_region_url_map" "foobar" {
	region          = "us-central1"
	name            = "regionurlmap-test-%s"
	default_service = "${google_compute_region_backend_service.foobar.self_link}"

	host_rule {
		hosts        = ["mysite.com", "myothersite.com"]
		path_matcher = "boop"
	}

	path_matcher {
		default_service = "${google_compute_region_backend_service.foobar.self_link}"
		name            = "boop"
	}

	test {
		host    = "mysite.com"
		path    = "/*"
		service = "${google_compute_region_backend_service.foobar.self_link}"
	}
}
`, randomSuffix, randomSuffix, randomSuffix)
}

func testAccComputeRegionUrlMap_ilb(randomSuffix string) string {
	return fmt.Sprintf(`
resource "google_compute_region_url_map" "foobar" {
  name        = "regionurlmap-test-%s"
  description = "a description"
  default_service = "${google_compute_region_backend_service.home.self_link}"

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name = "allpaths"
    default_service = "${google_compute_region_backend_service.home.self_link}"

    path_rule {
      paths   = ["/home"]
      route_action {
        cors_policy {
          allow_credentials = true
          allow_headers = ["Allowed content"]
          allow_methods = ["GET"]
          allow_origins = ["Allowed origin"]
          expose_headers = ["Exposed header"]
          max_age = 30
        }
        fault_injection_policy {
          abort {
            http_status = 234
            percentage = 5.6
          }
          delay {
            fixed_delay {
              seconds = 0
              nanos = 50000
            }
            percentage = 7.8
          }
        }
        request_mirror_policy {
          backend_service = "${google_compute_region_backend_service.home.self_link}"
        }
        retry_policy {
          num_retries = 4
          per_try_timeout {
            seconds = 30
          }
          retry_conditions = ["5xx", "deadline-exceeded"]
        }
        timeout {
          seconds = 20
          nanos = 750000000
        }
        url_rewrite {
          host_rewrite = "A replacement header"
          path_prefix_rewrite = "A replacement path"
        }
        weighted_backend_services {
          backend_service = "${google_compute_region_backend_service.home.self_link}"
          weight = 400
          header_action {
            request_headers_to_remove = ["RemoveMe"]
            request_headers_to_add {
              header_name = "AddMe"
              header_value = "MyValue"
              replace = true
            }
            response_headers_to_remove = ["RemoveMe"]
            response_headers_to_add {
              header_name = "AddMe"
              header_value = "MyValue"
            }
          }
        }
      }
    }
  }

  test {
    service = "${google_compute_region_backend_service.home.self_link}"
    host    = "hi.com"
    path    = "/home"
  }
}

resource "google_compute_region_backend_service" "home" {
  name        = "regionurlmap-test-%s"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = ["${google_compute_region_health_check.default.self_link}"]
  load_balancing_scheme = "INTERNAL_MANAGED"
}

resource "google_compute_region_backend_service" "home2" {
  name        = "regionurlmap-test2-%s"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = ["${google_compute_region_health_check.default.self_link}"]
  load_balancing_scheme = "INTERNAL_MANAGED"
}

resource "google_compute_region_health_check" "default" {
  name               = "regionurlmap-test-%s"
  http_health_check {}
}

`, randomSuffix, randomSuffix, randomSuffix, randomSuffix)
}
func testAccComputeRegionUrlMap_ilbUpdate(randomSuffix string) string {
	return fmt.Sprintf(`
resource "google_compute_region_url_map" "foobar" {
  name        = "regionurlmap-test-%s"
  description = "a description"
  default_service = "${google_compute_region_backend_service.home2.self_link}"

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name = "allpaths"
    default_service = "${google_compute_region_backend_service.home2.self_link}"

    path_rule {
      paths   = ["/home2"]
      route_action {
        cors_policy {
          allow_credentials = false
          allow_headers = ["Updated Allowed content"]
          allow_methods = ["GET", "POST"]
          allow_origins = ["Allowed origin", "Something else"]
          expose_headers = ["Exposed header", "Foo"]
          max_age = 31
        }
        fault_injection_policy {
          abort {
            http_status = 345
            percentage = 6.7
          }
          delay {
            fixed_delay {
              seconds = 1
              nanos = 40000
            }
            percentage = 9.1
          }
        }
        request_mirror_policy {
          backend_service = "${google_compute_region_backend_service.home.self_link}"
        }
        retry_policy {
          num_retries = 5
          per_try_timeout {
            seconds = 10
          }
          retry_conditions = ["5xx"]
        }
        timeout {
          seconds = 40
          nanos = 250000000
        }
        url_rewrite {
          host_rewrite = "A replacement"
          path_prefix_rewrite = "A replacement"
        }
        weighted_backend_services {
          backend_service = "${google_compute_region_backend_service.home.self_link}"
          weight = 300
          header_action {
            request_headers_to_remove = ["RemoveMe", "AndMe"]
            request_headers_to_add {
              header_name = "AddMe2"
              header_value = "MyValue2"
              replace = false
            }
            response_headers_to_remove = ["RemoveMe", "AndMe2"]
            response_headers_to_add {
              header_name = "AddMe2"
              header_value = "MyValue2"
            }
          }
        }
      }
    }
  }

  test {
    service = "${google_compute_region_backend_service.home.self_link}"
    host    = "hi.com"
    path    = "/home"
  }
}

resource "google_compute_region_backend_service" "home" {
  name        = "regionurlmap-test-%s"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = ["${google_compute_region_health_check.default.self_link}"]
  load_balancing_scheme = "INTERNAL_MANAGED"
}

resource "google_compute_region_backend_service" "home2" {
  name        = "regionurlmap-test2-%s"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = ["${google_compute_region_health_check.default.self_link}"]
  load_balancing_scheme = "INTERNAL_MANAGED"
}

resource "google_compute_region_health_check" "default" {
  name               = "regionurlmap-test-%s"
  http_health_check {}
}

`, randomSuffix, randomSuffix, randomSuffix, randomSuffix)
}
