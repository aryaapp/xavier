class CreateQuestionaires < ActiveRecord::Migration
  def change
    create_table :questionaires do |t|
      t.column :uuid, :uuid, null: false, default: "uuid_generate_v4()"
      t.string :title, :null => false     
      t.integer :questions, array: true, default: [], :null => false

      t.timestamps
    end
     
    change_column :questionaires, :uuid, :uuid, :null => false
    add_index :questionaires, :uuid
  end
end
