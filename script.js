// Mobile nav toggle
function toggleNav() {
  document.body.classList.toggle("nav-open");
  const nav = document.getElementById("nav");
  nav.classList.toggle("nav-open");
}

// Dynamic year in footer
document.getElementById("year").textContent = new Date().getFullYear();