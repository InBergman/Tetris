<template>
	<div id="app">
		<HelloWorld msg="Tetris"/>
	</div>
</template>

<script>

import HelloWorld from './components/HelloWorld.vue'
export default {
	name: 'app',
	data: function()
	{
		return {
			dropCounter		: 0,
			dropInterval	: 1000,
			lastTime		: 0,
			arena			: this.createMatrice(12, 20),
			player			: {},
			color			: [null, 'HotPink', 'Yellow', 'Lime', 'purple', 'DeepSkyBlue', 'Aqua', 'SpringGreen'],
			matrice			: [[]]
		}
	},
	components: { HelloWorld },
	mounted()
	{
		this.init();
		this.update();
	},
	methods:
	{
		init: function()
		{
			this.canvas   = document.getElementById('canvastetris');
			this.context  = this.canvas.getContext('2d');
			this.player   =  {
								pos: {x : 0, y : 0},
								matrice: this.matrice,
								score: 0
							};
			this.playerReset();
			this.context.scale(20, 20);
			this.context.fillStyle = '#000';
			this.context.fillRect(0, 0, this.canvas.width, this.canvas.height);
			document.addEventListener("keydown", event => {
				if(event.keyCode === 37)
					this.playerMove(-1);
				else if(event.keyCode === 39)
					this.playerMove(1);
				else if(event.keyCode === 40)
					this.playerDrop();
				else if(event.keyCode === 81)
					this.playerRotate(-1);
				else if(event.keyCode === 87)
					this.playerRotate( 1);
					
			})
		},
		
		playerReset: function()
		{
			const piece = 'TLZJSOI';
			this.player.matrice = this.createPiece(piece[piece.length * Math.random() | 0]);
			this.player.pos.y = 0;
			this.player.pos.x = (this.arena[0].length / 2 | 0) - (this.player.matrice[0].length / 2 | 0);
			if(this.collider())
			{
				this.arena.forEach(row => row.fill(0));
				this.player.score = 0;
				this.updateScore();
			}
		},

		createPiece: function(type)
		{
			if(type === 'T') 
			{
				return [[0, 0, 0],
						[1, 1, 1],
						[0, 1, 0]] 
			}
			else if(type === 'I') 
			{
				return [[0, 2, 0, 0],
						[0, 2, 0, 0],
						[0, 2, 0, 0],
						[0, 2, 0, 0]] 
			}
			else if(type === 'S') 
			{
				return [[0, 3, 3],
						[3, 3, 0],
						[0, 0, 0]] 
			}
			else if(type === 'Z') 
			{
				return [[4, 4, 0],
						[0, 4, 4],
						[0, 0, 0]] 
			}
			else if(type === 'O') 
			{
				return [[5, 5],
						[5, 5]] 
			}
			else if(type === 'L') 
			{
				return [[0, 6, 0],
						[0, 6, 0],
						[0, 6, 6]] 
			}
			else if(type === 'J') 
			{
				return [[0, 0, 0],
						[0, 7, 0],
						[7, 7, 7]] 
			}
		},

		arenaSweep: function()
		{
			let y = this.arena.length;	
			let count = 1;
			outer: for(y = this.arena.length - 1; y > 0; --y) 
					{
						for(let x = 0; x < this.arena[y].length; ++x)
						{
							if(this.arena[y][x] === 0)
								continue outer;
						}
						const newRow = this.arena.splice(y, 1)[0].fill(0);
						this.arena.unshift(newRow);
						this.player.score += count * 10; 
						count *= 2;
						++y;
					}
		},

		updateScore: function()
		{
			 document.getElementById("score").innerText = this.player.score;
		},

		playerRotate: function(dir)
		{
			const pos = this.player.pos.x;
			let offset = 1;
			this.rotate(this.player.matrice, dir);
			while(this.collider())
			{
				this.player.pos.x += offset;
				offset = -(offset + (offset > 0 ? 1 : -1));
				if(offset > this.player.matrice[0].length)
				{
					this.rotate(this.player.matrice, -dir);
					this.player.pos.x = pos;
					return;
				}
			}
		},

		rotate: function(matrice, dir)
		{
			for (let y = 0; y < matrice.length; ++y)
			{	
				for (let x = 0; x < y; ++x)
				{
					[ matrice[x][y], matrice[y][x] ] = [ matrice[y][x], matrice[x][y] ]
				}
			}
			if(dir > 0)
				matrice.forEach(row => row.reverse());
			else
				matrice.reverse();
		},

		playerMove: function(dir)
		{
			this.player.pos.x += dir;
			if(this.collider())
				this.player.pos.x -= dir;
		},

		collider: function()
		{
			const [mat, pos] = [this.player.matrice, this.player.pos];

			for (let y = 0; y < mat.length; ++y)
			{	
				for (let x = 0; x < mat[y].length; ++x)
				{
					if (mat[y][x] !== 0 &&  (this.arena[y + pos.y] && this.arena[y + pos.y][x + pos.x]) !== 0) 
					{
						return true
					}
				}	
			}
			return false
		},

		playerDrop: function()
		{
			this.player.pos.y++;
			if (this.collider())
			{
				this.player.pos.y--;
				this.merge();
				this.playerReset();
				this.arenaSweep();
				//this.updateScore();
			}
			this.dropCounter = 0;	
		},

		merge: function()
		{
			this.player.matrice.forEach((row, y) => {
				row.forEach((value, x) => {
					if(value !== 0)
					{
						this.arena[y + this.player.pos.y][x + this.player.pos.x] = value;
					}
				})
			})
		},

		createMatrice: function(w, h)
		{
			const newMatrice = [];
			while(h--)
				newMatrice.push(new Array(w).fill(0));
			return newMatrice;
		},

		draw: function()
		{
			this.context.fillStyle = '#000';
			this.context.fillRect(0, 0, this.canvas.width, this.canvas.height);
			this.drawMatrice(this.arena, {x:0, y:0});
			this.drawMatrice(this.player.matrice, this.player.pos);
		},

		update: function(time = 0)
		{
			const deltaTime = time - this.lastTime;

			this.lastTime = time;
			this.dropCounter += deltaTime;
			if(this.dropCounter > this.dropInterval)
			{
				this.playerDrop();
			}
			this.draw();
			this.updateScore();
			requestAnimationFrame(this.update)
		},

		drawMatrice: function(matrice, offset)
		{
			matrice.forEach((row, y) => {
				row.forEach((value, x) => {
					if(value !== 0)
					{
						this.context.fillStyle = this.color[value];
						this.context.fillRect(x + offset.x,
												y + offset.y,
			    								1, 1);
					}
				})
		})
		}
	}
}
</script>

<style>

#app {
	font-family: 'Avenir', Helvetica, Arial, sans-serif;
	-webkit-font-smoothing: antialiased;
	-moz-osx-font-smoothing: grayscale;
	text-align: center;
	color: #2c3e50;
	margin-top: 60px;
}
</style>
