import { objectNames } from "@wailsio/runtime";
import type { SignalDiagram } from "../../bindings/github.com/chrisjm66/openrcs/internal/signal_diagram";

const SIGNAL_RADIUS = 5
const TRACK_STROKE = 5

export function drawTracks(diagram: SignalDiagram | undefined, context: CanvasRenderingContext2D) {
if (!diagram || !diagram.tracks) {
				console.log('null');
				return;
			} else {
				Object.entries(diagram.tracks).forEach(([id, track]) => {
					console.log(id);
					if (track && track.diagramPositions) {
						console.log(track);

						for (let i = 1; i < track.diagramPositions.length; i++) {
							const previousPosition = track.diagramPositions[i - 1];
							const currentPosition = track.diagramPositions[i];
							context.beginPath();
							context.moveTo(previousPosition.x, previousPosition.y);
							context.lineTo(currentPosition.x, currentPosition.y);
							context.lineWidth = TRACK_STROKE;
							context.strokeStyle = 'darkgrey';
							context.stroke();
							// TODO perform state lookup to add headcode, color, etc.
						}
					} else {
						console.log('circuit is null');
					}
				});
			}
        }


export function drawSignals(diagram: SignalDiagram | undefined, context: CanvasRenderingContext2D) {
    if (!diagram || !diagram.signals) {
        return
    }

    Object.entries(diagram.signals).forEach(([id, signal]) => {
        if(signal && signal.diagramPosition) {
			context.beginPath()
            context.arc(signal.diagramPosition.x, signal.diagramPosition.y, SIGNAL_RADIUS, 0, 360, false)
			context.fillStyle = 'yellow'
			context.fill()
        }
    })
}