apiVersion: v1
kind: ServiceAccount
metadata:
  name: spire-agent
  namespace: {{ .Release.Namespace }}
