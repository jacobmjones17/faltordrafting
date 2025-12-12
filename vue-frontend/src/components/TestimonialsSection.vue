<template>
  <section id="testimonials" class="testimonials-section">
    <div class="container">
      <div class="section-heading">
        <div class="section-eyebrow">Testimonials</div>
        <h2>What clients and builders are saying.</h2>
        <p>
          These are placeholder testimonials — swap them out with your own client feedback as projects are completed.
        </p>
      </div>
      <div class="testimonials-grid">
        <div v-for="testimonial in testimonials" :key="testimonial.id" class="testimonial-card">
          <p class="testimonial-text">{{ testimonial.text }}</p>
          <div class="testimonial-meta">
            <div>
              <div class="testimonial-name">{{ testimonial.name }}</div>
              <div>{{ testimonial.location }}</div>
            </div>
            <div class="testimonial-badge">{{ testimonial.badge }}</div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  name: 'TestimonialsSection',
  data() {
    return {
      testimonials: [
        {
          id: 1,
          text: 'The plans were clear, detailed, and easy to build from. Our framing crew had everything they needed, which kept the project moving on schedule.',
          name: 'Local Builder',
          location: 'Wayne County, NC',
          badge: 'Custom Home'
        },
        {
          id: 2,
          text: 'We started with a rough sketch on paper and ended up with a full set of drawings that made sense to our contractor and the inspector. The process was smooth and efficient.',
          name: 'Homeowner',
          location: 'Eastern NC',
          badge: 'Remodel & Addition'
        },
        {
          id: 3,
          text: 'The 3D views helped us visualize the exterior and make decisions on windows and rooflines before we ever ordered materials.',
          name: 'Homeowner',
          location: 'Goldsboro, NC',
          badge: '3D Renderings'
        }
      ]
    }
  }
}
</script>