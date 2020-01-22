// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccSecretmanagerSecretIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/viewer",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccSecretmanagerSecretIamBinding_basicGenerated(context),
			},
			{
				// Test Iam Binding update
				Config: testAccSecretmanagerSecretIamBinding_updateGenerated(context),
			},
		},
	})
}

func TestAccSecretmanagerSecretIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/viewer",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccSecretmanagerSecretIamMember_basicGenerated(context),
			},
		},
	})
}

func TestAccSecretmanagerSecretIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/viewer",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccSecretmanagerSecretIamPolicy_basicGenerated(context),
			},
			{
				Config: testAccSecretmanagerSecretIamPolicy_emptyBinding(context),
			},
		},
	})
}

func testAccSecretmanagerSecretIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_secretmanager_secret" "secret-basic" {
  provider = google-beta

  secret_id = "tf-test-secret-%{random_suffix}"
  
  labels = {
    label = "my-label"
  }

  replication {
    user_managed {
      replicas {
        location = "us-central1"
      }
      replicas {
        location = "us-east1"
      }
    }
  }
}

resource "google_secretmanager_secret_iam_member" "foo" {
  provider = google-beta
  project = google_secretmanager_secret.secret-basic.project
  secret_id = google_secretmanager_secret.secret-basic.secret_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccSecretmanagerSecretIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_secretmanager_secret" "secret-basic" {
  provider = google-beta

  secret_id = "tf-test-secret-%{random_suffix}"
  
  labels = {
    label = "my-label"
  }

  replication {
    user_managed {
      replicas {
        location = "us-central1"
      }
      replicas {
        location = "us-east1"
      }
    }
  }
}

data "google_iam_policy" "foo" {
  provider = google-beta
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_secretmanager_secret_iam_policy" "foo" {
  provider = google-beta
  project = google_secretmanager_secret.secret-basic.project
  secret_id = google_secretmanager_secret.secret-basic.secret_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccSecretmanagerSecretIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_secretmanager_secret" "secret-basic" {
  provider = google-beta

  secret_id = "tf-test-secret-%{random_suffix}"
  
  labels = {
    label = "my-label"
  }

  replication {
    user_managed {
      replicas {
        location = "us-central1"
      }
      replicas {
        location = "us-east1"
      }
    }
  }
}

data "google_iam_policy" "foo" {
  provider = google-beta
}

resource "google_secretmanager_secret_iam_policy" "foo" {
  provider = google-beta
  project = google_secretmanager_secret.secret-basic.project
  secret_id = google_secretmanager_secret.secret-basic.secret_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccSecretmanagerSecretIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_secretmanager_secret" "secret-basic" {
  provider = google-beta

  secret_id = "tf-test-secret-%{random_suffix}"
  
  labels = {
    label = "my-label"
  }

  replication {
    user_managed {
      replicas {
        location = "us-central1"
      }
      replicas {
        location = "us-east1"
      }
    }
  }
}

resource "google_secretmanager_secret_iam_binding" "foo" {
 
  provider = google-beta
  project = google_secretmanager_secret.secret-basic.project
  secret_id = google_secretmanager_secret.secret-basic.secret_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccSecretmanagerSecretIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_secretmanager_secret" "secret-basic" {
  provider = google-beta

  secret_id = "tf-test-secret-%{random_suffix}"
  
  labels = {
    label = "my-label"
  }

  replication {
    user_managed {
      replicas {
        location = "us-central1"
      }
      replicas {
        location = "us-east1"
      }
    }
  }
}

resource "google_secretmanager_secret_iam_binding" "foo" {
  provider = google-beta
  project = google_secretmanager_secret.secret-basic.project
  secret_id = google_secretmanager_secret.secret-basic.secret_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:paddy@hashicorp.com"]
}
`, context)
}
