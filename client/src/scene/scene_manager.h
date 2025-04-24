#include "../window/window.h"
class SceneManager {
	public:
	// Method
	virtual void DrawScene();
};

class GameScene : public SceneManager {
	public:
	GameScene();
	~GameScene();
	void DrawScene() override;
};

enum SceneState {
	MENU,
	GAME,
};
