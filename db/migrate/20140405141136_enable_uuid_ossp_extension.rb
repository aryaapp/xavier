class EnableUuidOsspExtension < ActiveRecord::Migration
  def change
    enable_extension 'uuid-ossp'
  end
  
  def down
    disable_extension 'uuid-ossp' 
  end
end
