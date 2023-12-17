<template>
  <div class="notification-list">
    <div v-if="notifications.length">
      <div v-for="notification in notifications" :key="notification.id" class="notification-item">
        <div class="notification-content">
          {{ notification.message }}
          <small class="text-muted time">{{ timeAgo(notification.createdAt) }}</small>
        </div>
      </div>
    </div>
    <div v-else>
      <div class="notification-item">Brak powiadomien</div>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  notifications: { id: number; message: string; read: boolean; createdAt: string }[]
}>()

function timeAgo(timestamp: string) {
  const now = new Date()
  const createdAt = new Date(timestamp)
  const diff = now.getTime() - createdAt.getTime()
  const seconds = Math.floor(diff / 1000)
  if (seconds < 60) {
    return `${seconds} seconds ago`
  } else if (seconds < 3600) {
    const minutes = Math.floor(seconds / 60)
    return `${minutes} minutes ago`
  } else if (seconds < 86400) {
    const hours = Math.floor(seconds / 3600)
    return `${hours} hours ago`
  } else {
    const days = Math.floor(seconds / 86400)
    return `${days} days ago`
  }
}
</script>

<style scoped lang="scss">
/* The same CSS styles as before */
.notifications {
  position: relative;
  display: inline-block;
  margin-right: 20px;
}

.notification-icon {
  cursor: pointer;
}

.notification-popover {
  // position: absolute;
  top: 100%;
  right: 0;
  background-color: white;
  border: 1px solid #ccc;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  z-index: 100100;
}

.notification-list {
  max-height: 200px;
  overflow-y: auto;
}

.notification-item {
  padding: 10px;
  border-bottom: 1px solid #eee;
}

.notification-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 12px;
}

.time {
  display: flex;
  padding-left: 10px;
  margin-left: 10px;
  border-left: 1px solid var(--bs-border-color);
}
</style>
