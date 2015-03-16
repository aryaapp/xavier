class CreateNotifications < ActiveRecord::Migration
  def change
    create_table :notifications do |t|
      t.column :uuid, :uuid, default: "uuid_generate_v4()", :null => false

      t.string  :message, :null => false
      t.boolean :read, default: false      
      
      t.integer :object_id
      t.string  :object_type
      t.string  :object_uri

      t.timestamps
    end

    change_column :notifications, :uuid, :uuid, :null => false
    add_index :notifications, :uuid

    add_reference :notifications, :app, index: true    
    add_reference :notifications, :user, index: true    
  end
end
