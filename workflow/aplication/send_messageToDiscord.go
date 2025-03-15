package aplication

import "fmt"

func SendMessageToDiscord(action, workflow, status, url, repository string) string {
	var message string
	if action == "" {
		action = "no se encontro"
	}
	if workflow == "" {
		workflow = "no se encontro"
	}
	if status == "" {
		status = "desconocido"
	}
	switch action {
	case "complete":
		if status == "success" {
			message = "se ha enviado el mensaje con exito"
		} else if status == "failed" {
			message = "no se pudo enviar el mensaje"
		} else {
			message = fmt.Sprintf("error con del workflow: %s", status)
		}
	case "requested":
		message = "workflow iniciado"
	default:
		message = fmt.Sprintf("Evento: %s", action)
	}
	return fmt.Sprintf("%s en el repositorio %s\nWorkflow: %s\nDetalles: %s\n\n---\n",
		message, repository, workflow, url)
}
