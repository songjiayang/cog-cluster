from cog import BasePredictor, BaseModel, Input, Path
from deepface import DeepFace

class Output(BaseModel):
    gender: str
 
class Predictor(BasePredictor):
    def setup(self):
      print("load model here if need")

    def predict(self,
        image: Path = Input(description="input image url")
    ) -> Output:
        objs = DeepFace.analyze(
          img_path = image,
          actions = ['gender'],
          detector_backend = "retinaface",
          expand_percentage = 20,
        )
        return Output(gender=objs[0]["dominant_gender"])
