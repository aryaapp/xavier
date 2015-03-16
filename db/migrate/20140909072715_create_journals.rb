class CreateJournals < ActiveRecord::Migration
  def change
    create_table :journals do |t|
      t.column :uuid, :uuid, default: "uuid_generate_v4()", :null => false
      t.float :feeling, default: -1, :null => false
      t.jsonb :questions, :null => false

      t.timestamps
    end
    
    change_column :journals, :uuid, :uuid, :null => false
    add_index :journals, :uuid
    add_reference :journals, :app, index: true
    add_reference :journals, :user, index: true
  end
end
