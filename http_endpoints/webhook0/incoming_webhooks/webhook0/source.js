exports = async function(arg){
  const { default: nearestPointOnLine } = require('@turf/nearest-point-on-line');

const point = {type: 'Point', coordinates: [-77.037076, 38.884017]};
const line = {type: 'LineString', coordinates: [
    [-77.031669, 38.878605],
    [-77.029609, 38.881946],
    [-77.020339, 38.884084],
    [-77.025661, 38.885821],
    [-77.021884, 38.889563],
    [-77.019824, 38.892368]
]};


  console.log('Start nearestPointOnLine: ', new Date().getMilliseconds())
  const pointSnapped = nearestPointOnLine(line,point )  
  console.log('Finish nearestPointOnLine: ', new Date().getMilliseconds())
  return pointSnapped;
};