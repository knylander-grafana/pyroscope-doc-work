require_relative "boot"

require "rails"
# Pick the frameworks you want:
require "active_model/railtie"
# require "active_job/railtie"
require "active_record/railtie"
# require "active_storage/engine"
require "action_controller/railtie"
# require "action_mailer/railtie"
# require "action_mailbox/engine"
# require "action_text/engine"
require "action_view/railtie"
# require "action_cable/engine"
require "rails/test_unit/railtie"
require "pyroscope"

# Require the gems listed in Gemfile, including any gems
# you've limited to :test, :development, or :production.
Bundler.require(*Rails.groups)

module RideshareRails
  class Application < Rails::Application
    # Initialize configuration defaults for originally generated Rails version.
    config.load_defaults 7.0

    # Configuration for the application, engines, and railties goes here.
    #
    # These settings can be overridden in specific environments using the files
    # in config/environments, which are processed later.
    #
    # config.time_zone = "Central Time (US & Canada)"
    # config.eager_load_paths << Rails.root.join("extras")

    # Don't generate system test files.
    config.generators.system_tests = nil
    config.action_controller.include_all_helpers = true
    config.active_record.sqlite3_production_warning=false

    Pyroscope.configure do |config|
      config.app_name       = ENV.fetch("PYROSCOPE_APPLICATION_NAME", "rails-ride-sharing-app")
      config.server_address = ENV.fetch("PYROSCOPE_SERVER_ADDRESS", "http://pyroscope:4040")
      config.auth_token     = ENV.fetch("PYROSCOPE_AUTH_TOKEN", "")

      config.tags = {
        "region": ENV["REGION"] || "us-east",
      }
    end
  end
end