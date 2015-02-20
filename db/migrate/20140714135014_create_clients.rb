class CreateClients < ActiveRecord::Migration
  def change
    create_table :clients do |t|
      t.column :uuid, :uuid, default: "uuid_generate_v4()", :null => false
      t.string :name, :null => false 
      t.string :url, :null => false
      t.string :secret, :null => false       
      t.string :grant_types, array: true, default: [], :null => false
      t.string :permitted_scopes, array: true, default: [], :null => false

      t.timestamps
    end
    change_column :clients, :uuid, :uuid, :null => false
    add_index :clients, :uuid
  end
end
