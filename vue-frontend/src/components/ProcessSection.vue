<template>
  <section id="process">
    <div class="container">
      <div class="section-heading">
        <div class="section-eyebrow">Process</div>
        <h2>A straightforward process from sketch to construction.</h2>
        <p>
          Clear expectations, consistent communication, and drawings tailored to how builders actually work in Eastern NC.
        </p>
      </div>
      <div class="process-grid">
        <div v-for="step in processSteps" :key="step.number" class="process-step">
          <div class="process-step-header">
            <div class="process-number">{{ step.number }}</div>
            <h3>{{ step.title }}</h3>
          </div>
          <p>{{ step.description }}</p>
          <p v-if="step.hint" class="hint">{{ step.hint }}</p>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  name: 'ProcessSection',
  data() {
    return {
      processSteps: [
        {
          number: 1,
          title: 'Discovery & Sketches',
          description: 'Share your ideas, inspiration, and any existing sketches or plans. We\'ll clarify your goals, budget, and timeline so we\'re designing for the real world.',
          hint: 'You can upload sketches directly in the quote form below.'
        },
        {
          number: 2,
          title: 'Drafting & Revisions',
          description: 'FALTOR Consulting prepares draft plans for review, then fine-tunes layouts and details based on your feedback and builder input.'
        },
        {
          number: 3,
          title: 'Final Plans & Permitting',
          description: 'You receive clean, permit-ready drawings and optional 3D views, ready for your contractor and local permitting office.'
        }
      ]
    }
  }
}
</script>