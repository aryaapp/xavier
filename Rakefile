require 'active_record'
require 'active_record_migrations'

ActiveRecordMigrations.configure do |c|
  c.database_configuration = {
    'development' => { 
      'adapter' => 'postgresql', 
      'database' => 'arya_development'
    },
   	'test' => { 
      'adapter' => 'postgresql', 
      'database' => 'arya_test'
    }
  }
  # c.db_dir = 'schema'
  # c.migrations_paths = ['schema/migrate']
end

ActiveRecordMigrations.load_tasks