<template>
  <section id="gallery">
    <div class="container">
      <div class="section-heading">
        <div class="section-eyebrow">3D Rendering Gallery</div>
        <h2>Visualize your home before you break ground.</h2>
        <p>
          High-quality 3D views and exterior renderings help you and your builder make confident decisions about
          layout, massing, and finishes.
        </p>
      </div>
      <div class="gallery-grid">
        <!-- Gallery items would go here - same as original HTML -->
        <div class="gallery-item" style="display:flex;flex-direction:column;gap:4px;min-height:200px;background:#f5f7fb;">
          <!-- <img src="/all sides plan.png" alt="Drafting Elevations" />
          <img src="/Render All Sides.png" alt="Rendered Elevations" /> -->
          <div style="padding:2rem;text-align:center;color:var(--text-muted);">
            <p>Gallery images will be added here</p>
          </div>
          <div class="gallery-caption">
            <span>Drafting + Rendering Comparison</span>
            <span>FALTOR Consulting</span>
          </div>
        </div>
        <!-- Add other gallery items as needed -->
      </div>
    </div>
  </section>
</template>

<script>
export default {
  name: 'GallerySection'
}
</script>