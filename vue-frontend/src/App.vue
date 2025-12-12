<template>
  <div id="app">
    <!-- Header Component -->
    <HeaderComponent />
    
    <!-- Hero Section -->
    <HeroSection />
    
    <!-- Services Section -->
    <ServicesSection />
    
    <!-- Gallery Section -->
    <GallerySection />
    
    <!-- Process Section -->
    <ProcessSection />
    
    <!-- Testimonials Section -->
    <TestimonialsSection />
    
    <!-- Contact Section -->
    <ContactSection />
    
    <!-- Footer Component -->
    <FooterComponent />
  </div>
</template>

<script>
import HeaderComponent from './components/HeaderComponent.vue'
import HeroSection from './components/HeroSection.vue'
import ServicesSection from './components/ServicesSection.vue'
import GallerySection from './components/GallerySection.vue'
import ProcessSection from './components/ProcessSection.vue'
import TestimonialsSection from './components/TestimonialsSection.vue'
import ContactSection from './components/ContactSection.vue'
import FooterComponent from './components/FooterComponent.vue'

export default {
  name: 'App',
  components: {
    HeaderComponent,
    HeroSection,
    ServicesSection,
    GallerySection,
    ProcessSection,
    TestimonialsSection,
    ContactSection,
    FooterComponent
  }
}
</script>