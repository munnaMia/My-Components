const icon = document.getElementById("containerID")
const ele = document.getElementById("mainContainer")

icon.addEventListener("click", ()=>{
    icon.classList.toggle("openContainer")
    ele.classList.toggle("click-container")
})