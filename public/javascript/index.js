function showBookmarkDeleteModal(bookmarkID) {
  let modal = document.getElementById("bookmark-delete-modal");
  if (!modal) {
    console.log("error: failed to find bookmark delete modal");
    return;
  }

  let form = document.getElementById("bookmark-delete-form");
  if (!form) {
    console.log("error: failed to find bookmark delete form");
    return;
  }

  form.action = "/bookmarks/"+bookmarkID+"/destroy";

  modal.style.display = "flex";
  return;
}

function hideBookmarkDeleteModal() {
  let modal = document.getElementById("bookmark-delete-modal");
  if (!modal) {
    console.log("error: failed to find bookmark delete modal");
    return;
  }

  let form = document.getElementById("bookmark-delete-form");
  if (!form) {
    console.log("error: failed to find bookmark delete form");
    return;
  }

  form.action = "/bookmarks/__ID__/destroy";

  modal.style.display = "none";
  return;
}
