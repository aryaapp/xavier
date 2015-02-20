class CreateRelations < ActiveRecord::Migration
  def change
    create_table :relations do |t|
      t.integer :observer_id, :null => false 
      t.integer :subject_id, :null => false 
      t.string :type, :null => false 
      t.string :permissions

      t.timestamps
    end
    
    add_index :relations, :observer_id
    add_index :relations, :subject_id
    add_index :relations, [:observer_id, :subject_id], unique: true
  end
end
