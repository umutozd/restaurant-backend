# Endpoints and Specifications

Currently, there are 11 elementary endpoints necessary for the system. Here are their specifications:

## Table of Contents
- [ListMenu](#listmenu)
- [CreateMenuItem](#createmenuitem)
- [ListMenuItems](#listmenuitems)
- [UpdateMenuItem](#updatemenuitem)
- [DeleteMenuItem](#deletemenuitem)
- [CreateCategory](#createcategory)
- [ListCategories](#listcategories)
- [UpdateCategory](#updatecategory)
- [DeleteCategory](#deletecategory)
- [UpdateCart](#updatecart)
- [UpdateCartItemsStatus](#updatecartitemsstatus)


### CreateMenuItem
- Description: Creates a MenuItem object.
  - HTTP:
    ```http
    POST /api/v1/menu/items/create
    ```
  - Body: `MenuItem` object.
  - Response: `MenuItem` object with its ID set.

### ListMenu
  - Description: Finds all categories, menu items and groups items by the categories they belong to.
  - HTTP:
    ```http
    GET /api/v1/menu/list
    ```
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

### ListMenuItems
  - Description: Lists all `MenuItem` objects.
  - HTTP:
    ```http
    GET /api/v1/menu/items/list
    ```
  - Body: None
  - Response:
    ```Typescript
    type MenuItemsList MenuItem[];
    ```

### UpdateMenuItem
  - Description: Updated a `MenuItem` specified by its ID.
  - HTTP:
    ```http
    POST /api/v1/menu/items/update
    ```
  - Body:
    ```Typescript
    interface UpdateMenuItemReq {
      /** Specifies which fields are to be updated */
      fields: UpdateMenuItemField[];
      /** Contains the id of the item to update and new values for specified fields */
      item: MenuItem;
    }
    ```
  - Response: The updated `MenuItem` object.

### DeleteMenuItem
  - Description: Deletes a `MenuItem` object, specified by its ID.
  - HTTP:
    ```http
    DELETE /api/v1/menu/items/delete
    ```
  - Body: 
    ```Typescript
    interface DeleteMenuItemReq {
      /** ID of the MenuItem to be deleted */
      id: string;
    }
    ```
  - Response: Unused.

### CreateCategory
  - Description: Creates a `Category` object.
  - HTTP:
    ```http
    POST /api/v1/category/create
    ```
  - Body: `Category` object
  - Response: `Category` object, with its id set.

### ListCategories
  - Description: Lists all `Category` objects.
  - HTTP:
    ```http
    GET /api/v1/category/list
    ```
  - Body: None
  - Response: 
    ```Typescript
    type CategoryList Category[];
    ```

### UpdateCategory
  - Description: Updates a `Category` object, specified by its ID.
  - HTTP:
    ```http
    POST /api/v1/category/update
    ```
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

### DeleteCategory
  - Description: Deletes a `Category` object, specified by its ID.
  - HTTP:
    ```http
    DELETE /api/v1/category/delete
    ```
  - Body: 
    ```Typescript
    interface DeleteCategoryReq {
      /** ID of the category to be deleted */
      id: string;
    }
    ```
  - Response: Unused.

### UpdateCart
  - Description: Adds/removes items to/from a `Cart` object. Each `CartItem` created will have `status=0`.
  - HTTP:
    ```http
    POST /api/v1/cart/update
    ```
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

### UpdateCartItemsStatus
  - Description: Updates the states of `CartItem`'s in the `Cart`, specified by their ID's.
  - HTTP:
    ```http
    POST /api/v1/cart/update/items/status
    ```
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

