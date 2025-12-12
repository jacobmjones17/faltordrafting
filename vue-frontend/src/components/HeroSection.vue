<template>
  <section class="hero">
    <div class="container hero-grid">
      <div>
        <div class="eyebrow">Residential Drafting • Eastern North Carolina</div>
        <h1>
          Precise <span>house plans & remodel drafting</span> built for real-world construction.
        </h1>
        <p class="hero-tagline">
          FALTOR Consulting delivers clear, buildable drawings for custom homes, additions, remodels, and more —
          so your contractor has exactly what they need on day one.
        </p>
        <div class="hero-meta">
          <span>✓ Custom House Plans</span>
          <span>✓ Remodels & Additions</span>
          <span>✓ Permit-Ready Drawings</span>
        </div>
        <div class="hero-actions">
          <button class="btn btn-primary" @click="scrollToSection('quote')">
            Start a Project
          </button>
          <button class="btn btn-outline" @click="scrollToSection('services')">
            View Drafting Services
          </button>
        </div>
        <p class="hero-note">
          Typical projects include full construction sets for new homes, additions, and remodels across Eastern NC.
        </p>
      </div>

      <aside class="hero-visual">
        <div class="hero-visual-inner">
          <div class="hero-visual-header">
            <div class="hero-visual-title">Sample Floor Plan Layout</div>
            <div class="hero-visual-pill">Drafted by FALTOR Consulting</div>
          </div>
          <div class="hero-plan-preview">
            <div class="hero-plan-card">
              <h3>3 Bed • 2 Bath Craftsman</h3>
              <div class="hero-plan-metadata">
                <span>~1,950 sq ft</span>
                <span>9' Ceilings</span>
                <span>Crawlspace</span>
              </div>
              <p>
                Clean, construction-ready drawings with dimensions, notes, and elevations tailored to local building practices.
              </p>
              <div class="hero-plan-footer">
                <strong>Ready for permit sets</strong>
                <span style="font-size:0.78rem;color:#9ca3af;">Example project</span>
              </div>
            </div>
            <ul class="hero-mini-list">
              <li>Custom house plans from scratch or from your sketches</li>
              <li>Remodel & addition plans that tie into existing structure</li>
              <li>Optional 3D views to help visualize before you build</li>
              <li>Deliverables tailored to your builder or municipality</li>
            </ul>
          </div>
        </div>
      </aside>
    </div>
  </section>
</template>

<script>
export default {
  name: 'HeroSection',
  methods: {
    scrollToSection(sectionId) {
      const element = document.getElementById(sectionId)
      if (element) {
        element.scrollIntoView({ behavior: 'smooth' })
      }
    }
  }
}
</script>