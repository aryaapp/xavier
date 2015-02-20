class CreateDevices < ActiveRecord::Migration
  def change
    create_table :devices do |t|
      t.string :token, :null => false
      t.string :environment, :null => false 
      t.string :name, :null => false 
      t.string :model, :null => false 
      t.string :os, :null => false 
      t.string :os_version, :null => false 
      t.string :app_version, :null => false 
      
      t.timestamps
    end
    add_index :devices, :token, unique: true
    add_reference :devices, :user, index: true    
  end
end
