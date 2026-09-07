package database

import "testing"

func TestConfigFromEnvUsesDefaults(t *testing.T) {
	for _, key := range []string{
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
		"DB_SSLMODE",
	} {
		t.Setenv(key, "")
	}

	cfg := configFromEnv()

	if cfg.host != "localhost" {
		t.Fatalf("expected default host localhost, got %q", cfg.host)
	}
	if cfg.port != "5432" {
		t.Fatalf("expected default port 5432, got %q", cfg.port)
	}
	if cfg.name != "go_generics" {
		t.Fatalf("expected default database go_generics, got %q", cfg.name)
	}
}

func TestConfigFromEnvUsesEnvironment(t *testing.T) {
	t.Setenv("DB_HOST", "database.internal")
	t.Setenv("DB_PORT", "6543")
	t.Setenv("DB_USER", "app_user")
	t.Setenv("DB_PASSWORD", "app_password")
	t.Setenv("DB_NAME", "app_database")
	t.Setenv("DB_SSLMODE", "require")

	cfg := configFromEnv()

	if cfg.host != "database.internal" || cfg.port != "6543" {
		t.Fatalf("unexpected address: %s:%s", cfg.host, cfg.port)
	}
	if cfg.user != "app_user" || cfg.password != "app_password" {
		t.Fatal("database credentials were not loaded from the environment")
	}
	if cfg.name != "app_database" || cfg.sslMode != "require" {
		t.Fatal("database settings were not loaded from the environment")
	}
}
