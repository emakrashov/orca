Orca can work with multiple stores.
Each store is represented as a single file.

working_dir/
  - orca_main_system_store.store
  - store1.store
  - store2.store

store can be of two types, system and user.

When the app starts, it the main store.
It loads other stores lazily.

If there are no stores, it prints the message informing the user that they are expected to create one.



---

