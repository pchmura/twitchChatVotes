import React, { Component } from 'react';
import {Bar as BarChart} from 'react-chartjs'
class LiveChart extends Component {
 
    render() {
        return (
            <div>
                <BarChart data={{
                labels: this.props.chartData.emotes,
                datasets: [{
                    label: "My First dataset",
                    fillColor: "rgba(220,220,220,0.5)",
                    strokeColor: "rgba(220,220,220,0.8)",
                    highlightFill: "rgba(220,220,220,0.75)",
                    highlightStroke: "rgba(220,220,220,1)",
                    data: this.props.chartData.data
                    }]
            }} />
            </div>
        );
    }
}

export default LiveChart;