# Types
In this document, all data types stored in the databases are documented. All type definitions are made in `TypeScript` so that they can be copied directly to the front-end repository.

## Table of Contents
- [Data Types](#data-types)
  - [Category](#category)
  - [MenuItem](#menuitem)
  - [CategoryAndItems](#categoryanditems)
  - [Menu](#menu)
- [Enum Types](#enum-types)
  - [CartItemStatus](#cartitemstatus)
  - [UpdateCategoryField](#updatecategoryfield)
  - [UpdateMenuItemField](#updatemenuitemfield)
- [Request Types](#request-types)
  - [DeleteCategoryReq](#deletecategoryreq)
  - [DeleteMenuItemReq](#deletemenuitemreq)
  - [UpdateCartItemsStatusReq](#updatecartitemsstatusreq)
  - [UpdateCartReq](#updatecartreq)
  - [UpdateCategoryReq](#updatecategoryreq)
  - [UpdateMenuItemReq](#updatemenuitemreq)

## Data Types
### Category
The menu is divided into categories (e.g. desserts, beverages etc.).
```Typescript
interface Cateogry {
  /** ID of the Category. */
  id: string;
  /** Name of the Category. */
  name: string;
}
```

### MenuItem
Menu items are the most basic elements of the menu. Each item must belong to a category and have a non-empty name.
```Typescript
interface MenuItem {
  /** ID of the Category. */
  id: string;
  /** ID of the category this MenuItem is in. Cannot be empty. */
  category_id: string;
  /** Name of the item. */
  name: string;
  price: number;
  /** URL of the image to be displayed in the front-end. */
  image: string;
}
```

### CategoryAndItems
This is an element of the upper-most view of the menu. It contains a category and all the items in that category.
```Typescript 
interface CategoryAndItems {
  category: Category;
  items: MenuItems[];
}
```

### Menu
The upper-most view of the menu.
```Typescript
type Menu = CategoryAndItems[];
```

## Enum Types
### CartItemStatus
```Typescript
type CartItemStatus = "ORDERED" | "APPROVED" | "PAID";

const CartItemStatus_ORDERED: CartItemStatus = "ORDERED";
const CartItemStatus_APPROVED: CartItemStatus = "APPROVED";
const CartItemStatus_PAID: CartItemStatus = "PAID";

const cartItemStatuses: { [key: string]: number } = {
  [CartItemStatus_ORDERED]: 0,
  [CartItemStatus_APPROVED]: 1,
  [CartItemStatus_PAID]: 2,
};
```

### UpdateCategoryField
```Typescript
type UpdateCategoryField = "CATEGORY_NAME";

const UpdateCategoryField_CATEGORY_NAME: UpdateCategoryField = "CATEGORY_NAME";

const updateCategoryFields: { [key: string]: number } = {
  [UpdateCategoryField_CATEGORY_NAME]: 0,
};
```

### UpdateMenuItemField
```Typescript
type UpdateMenuItemField = "CATEGORY_ID" | "ITEM_NAME" | "PRICE" | "IMAGE";

const UpdateMenuItemField_CATEGORY_ID: UpdateMenuItemField = "CATEGORY_ID";
const UpdateMenuItemField_ITEM_NAME: UpdateMenuItemField = "ITEM_NAME";
const UpdateMenuItemField_PRICE: UpdateMenuItemField = "PRICE";
const UpdateMenuItemField_IMAGE: UpdateMenuItemField = "IMAGE";

const updateMenuItemFields: { [key: string]: number } = {
  [UpdateMenuItemField_CATEGORY_ID]: 0,
  [UpdateMenuItemField_ITEM_NAME]: 1,
  [UpdateMenuItemField_PRICE]: 2,
  [UpdateMenuItemField_IMAGE]: 3
};
```

## Request Types

### DeleteCategoryReq
```Typescript
interface DeleteCategoryReq {
  id: string;
}
```

### DeleteMenuItemReq
```Typescript
interface DeleteMenuItemReq {
  id: string;
}
```

### UpdateCartItemsStatusReq
```Typescript
interface UpdateCartItemsStatusReq {
  cart_id: string;
  item_ids: string[];
  status: CartItemStatus
}
```

### UpdateCartReq
```Typescript
interface UpdateCartReq {
  /** ID of the cart to update */
  id: string;
  /** MenuItem ID's to be added */
  add: string[];
  /** CarItem ID's to be removed */
  remove: string[];
}
```

### UpdateCategoryReq
```Typescript
interface UpdateCategoryReq {
  fields: UpdateCategoryField[];
  /** Category must contain ID */
  category: Category;
}
```

### UpdateMenuItemReq
```Typescript
interface UpdateMenuItemReq {
  fields: UpdateMenuItemFields[];
  /** MenuItem must contain ID */
  menu_item: MenuItem;
}
```