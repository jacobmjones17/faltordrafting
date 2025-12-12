<template>
  <header>
    <div class="container">
      <div class="nav" :class="{ 'nav-open': navOpen }" id="nav">
        <div class="nav-left">
          <!-- <img src="/logo.png" alt="FALTOR Consulting logo" class="nav-logo" /> -->
          <div class="nav-brand-text">
            <span>FALTOR Consulting</span>
            <span>Residential Drafting • Eastern NC</span>
          </div>
        </div>
        <nav class="nav-links">
          <a href="#services" @click="scrollToSection('services')">Services</a>
          <a href="#gallery" @click="scrollToSection('gallery')">3D Renderings</a>
          <a href="#process" @click="scrollToSection('process')">Process</a>
          <a href="#testimonials" @click="scrollToSection('testimonials')">Testimonials</a>
          <a href="#contact" @click="scrollToSection('contact')">Contact</a>
        </nav>
        <div class="nav-cta">
          <button class="btn btn-outline" @click="scrollToSection('contact')">
            Contact
          </button>
          <button class="btn btn-primary" @click="scrollToSection('quote')">
            Request a Quote
          </button>
        </div>
        <button class="nav-toggle" aria-label="Toggle navigation" @click="toggleNav">☰</button>
      </div>
    </div>
  </header>
</template>

<script>
export default {
  name: 'HeaderComponent',
  data() {
    return {
      navOpen: false
    }
  },
  methods: {
    toggleNav() {
      this.navOpen = !this.navOpen
      document.body.classList.toggle("nav-open", this.navOpen)
    },
    scrollToSection(sectionId) {
      const element = document.getElementById(sectionId)
      if (element) {
        element.scrollIntoView({ behavior: 'smooth' })
      }
      this.navOpen = false
      document.body.classList.remove("nav-open")
    }
  }
}
</script>