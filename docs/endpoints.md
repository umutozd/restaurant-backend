## Endpoints and Specifications

Currently, there are 11 elementary endpoints necessary for the system. Here are their specifications:

- **/api/v1/menu/list**:
  - Description: Finds all categories, menu items and groups items by the categories they belong to.
  - Method: `GET`
  - Body: None
  - Response:
    ```Typescript
    interface Menu {
      /** All categories */
      category: Category;
      /** All menu items */
      items: MenuItem[];
    }
    ```

- **/api/v1/menu/items/list**
  - Description: Lists all MenuItem objects.
  - Method: `GET`
  - Body: None
  - Response:
    ```Typescript
    type MenuItemsList MenuItem[];
    ```

- **/api/v1/menu/items/update**
  - Description: Updated a MenuItem specified by its id.
  - Method: `POST`
  - Body:
    ```Typescript
    interface UpdateMenuItemReq {
      /** Specifies which fields are to be updated */
      fields: UpdateMenuItemField[];
      /** Contains the id of the item to update and new values for specified fields */
      item: MenuItem;
    }
    ```
  - Response: The updated `MenuItem`

- **/api/v1/menu/items/delete**
  - Description: Deletes a MenuItem object.
  - Method: `DELETE`
  - Body: 
    ```Typescript
    interface DeleteMenuItemReq {
      /** ID of the MenuItem to be deleted */
      id: string;
    }
    ```
  - Response: Unused.

- **/api/v1/category/create**
  - Description: Creates a MenuItem object.
  - Method: `POST`
  - Body: `Category` object
  - Response: `Category` object, with its id set.

- **/api/v1/category/list**
  - Description: Lists all Category objects.
  - Method: `GET`
  - Body: None
  - Response: 
    ```Typescript
    type CategoryList Category[];
    ```

- **/api/v1/category/update**
  - Description: Updates a Category object, specified by its id.
  - Method: `POST`
  - Body: 
    ```Typescript
    interface UpdateCategoryReq {
      /** Specifies which fields are to be updated */
      fields: UpdateCategoryField[];
      /** Contains the id of the category to update and
        * new values for specified fields
        */
      category: Category;
    }
    ```
  - Response: The updated `Category` object.

- **/api/v1/category/delete**
  - Description: Deletes a Category object, specified by its id.
  - Method: `DELETE`
  - Body: 
    ```Typescript
    interface DeleteCategoryReq {
      /** ID of the category to be deleted */
      id: string;
    }
    ```
  - Response: Unused.

- **/api/v1/cart/update**
  - Description: Adds/removes items to/from a `Cart` object. Each CartItem created will have `status=0`.
  - Method: `POST`
  - Body: 
    ```Typescript
    interface UpdateCartReq {
      /** ID of the cart to update. If left empty, a new cart will be created. */
      id: string;
      /** ID's of the MenuItem objects to be added. */
      add: string[];
      /** CartItem ID's to remove from this cart. */
      remove: string[];
    }
    ```
  - Response: Updated `Cart` object.

- **/api/v1/cart/update/items/status**
  - Description: Updates the states of CartItem's in the Cart, specified by their ID's.
  - Method: `POST`
  - Body:
  ```Typescript
  interface UpdateCartItemsStatusReq {
    /** ID of the cart to be updated */
    cart_id: string;
    /** ID's of the cart items whose status' will be updated */
    item_ids: string[];
    /** The new status value */
    status: CartItemStatus;
  }
  ```
  - Response: Updated `Cart` object.

