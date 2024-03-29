class CreateApps < ActiveRecord::Migration
  def change
    create_table :apps do |t|
      t.column :uuid, :uuid, default: "uuid_generate_v4()", :null => false
      t.string :name, :null => false 
      t.string :url, :null => false
      t.string :secret, :null => false       
      t.string :grant_types, array: true, default: [], :null => false
      t.string :permitted_scopes, array: true, default: [], :null => false

      t.timestamps
    end
    change_column :apps, :uuid, :uuid, :null => false
    add_index :apps, :uuid
    add_reference :users, :app, index: true
  end
end
