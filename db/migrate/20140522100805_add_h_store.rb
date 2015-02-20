class AddHStore < ActiveRecord::Migration
  def change
    enable_extension 'hstore' 
  end
  
  def down
    disable_extension 'hstore' 
  end
end
